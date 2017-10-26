package tech.evove.goandroid;

import java.util.concurrent.TimeUnit;

import core.Core;
import core.ScheduleWorker;
import io.reactivex.Scheduler;
import io.reactivex.disposables.Disposable;
import io.reactivex.internal.util.ExceptionHelper;

public class GoScheduler extends Scheduler {
    private final core.Scheduler scheduler = Core.goScheduler();

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
