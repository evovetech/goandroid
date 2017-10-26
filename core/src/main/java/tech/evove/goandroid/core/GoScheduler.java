package tech.evove.goandroid.core;

import java.util.concurrent.TimeUnit;

import core.Core;
import core.ScheduleWorker;
import io.reactivex.Scheduler;
import io.reactivex.disposables.Disposable;
import io.reactivex.internal.util.ExceptionHelper;

public final class GoScheduler extends Scheduler {
    private static final class Holder {
        private static final GoScheduler Instance = new GoScheduler(Core.goScheduler());
    }

    private final core.Scheduler scheduler;

    private GoScheduler(core.Scheduler scheduler) {
        this.scheduler = scheduler;
    }

    public static Scheduler instance() {
        return Holder.Instance;
    }

    @Override
    public Worker createWorker() {
        return new GoWorker(scheduler.createWorker());
    }

    @Override
    public Disposable scheduleDirect(Runnable run, long delay, TimeUnit unit) {
        return schedule(scheduler, run, delay, unit);
    }

    static Disposable schedule(ScheduleWorker worker, Runnable run, long delay, TimeUnit unit) {
        try {
            core.Disposable disposable = worker.schedule(GoRunnable.wrap(run), unit.toNanos(delay));
            return GoDisposable.wrap(disposable);
        } catch (Exception e) {
            throw ExceptionHelper.wrapOrThrow(e);
        }
    }
}
